
<template>
  <main>
    <div style="display: block; margin: 0 auto">
      <el-col>
        <el-card shadow="hover">
          <el-button type="primary" @click="GenerateKeyPair">Generate</el-button>
          <div style="margin-top: 20px">
          </div>
          <el-input size="large" maxlength="128" show-word-limit v-model="PrivateKey" rows="3" type="textarea"
            placeholder="Private Key" />

          <div style="margin-top: 20px">
          </div>
          <el-input size="large" maxlength="128" show-word-limit v-model="PubKey" rows=3 type="textarea"
            placeholder="Public Key" />
        </el-card>
        <el-card shadow="hover" style="margin-top: 15px">
          <el-button type="primary" @click="SignAction">Sign</el-button>&nbsp;&nbsp;
          <el-button type="primary" @click="SignVerifyAction">Verify</el-button>
          <div style="margin-top: 20px">
          </div>
          <el-input size="large" maxlength="128" show-word-limit v-model="SignContent" rows="3" type="textarea"
            placeholder="Sign Content" spellcheck="false" />
          <div style="margin-top: 20px">
          </div>
          <el-input size="large" maxlength="128" show-word-limit v-model="Signature" rows="3" type="textarea"
            placeholder="Signature" />
        </el-card>
      </el-col>
    </div>
  </main>
</template>
  
<script lang="ts" setup>
import { ElMessage } from 'element-plus';
import { ref } from 'vue'
import { Ed25519Verify } from '../../wailsjs/go/main/App';
import { CreateEd25519KeyPair, Ed25519Sign } from '../service';

const PrivateKey = ref("");
const PubKey = ref("");
const SignContent = ref("");
const Signature = ref("");

const GenerateKeyPair = async () => {
  let KeyPair = await CreateEd25519KeyPair("hex");
  let PrivateStr = KeyPair.PrivateKey
  let PublicStr = KeyPair.PublicKey

  console.log(PrivateStr);
  console.log(PublicStr);

  PrivateKey.value = PrivateStr;
  PubKey.value = PublicStr;
}

const SignAction = async () => {
  let resp = await Ed25519Sign(PrivateKey.value, PubKey.value, SignContent.value, "base64");
  if (resp.errmsg != "") {
    Signature.value = resp.errmsg
  } else {
    Signature.value = resp.signature
  }
}

const SignVerifyAction = async () => {
  let resp = await Ed25519Verify(PubKey.value, SignContent.value, Signature.value);
  if (resp.errmsg != "") {
    ElMessage.error(resp.errmsg)
  } else {
    if (resp.isvalid) {
      ElMessage({
        message: 'Valid Signature',
        type: 'success',
      })
    } else {
      ElMessage({
        message: 'Invalid Signature',
        type: 'warning',
      })
    }

  }
}
</script>
