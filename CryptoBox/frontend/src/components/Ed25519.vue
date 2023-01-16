
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
      </el-col>
    </div>
  </main>
</template>
  
<script lang="ts" setup>
import { ref } from 'vue'
import { CreateEd25519KeyPair } from '../service';



const PrivateKey = ref("");
const PubKey = ref("");

const GenerateKeyPair = async () => {
  let KeyPair = await CreateEd25519KeyPair("hex");
  let PrivateStr = KeyPair.PrivateKey
  let PublicStr = KeyPair.PublicKey

  console.log(PrivateStr);
  console.log(PublicStr);

  PrivateKey.value = PrivateStr;
  PubKey.value = PublicStr;
}



</script>
