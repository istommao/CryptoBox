
<template>
    <main>
        <div style="display: block; margin: 0 auto">
            <el-col>
                <el-card shadow="hover">
                    <el-button type="primary" @click="GenerateRSAKeyResult">Generate</el-button>
                    <div style="margin-top: 20px">
                    </div>
                    <el-input size="large" maxlength="3400" show-word-limit v-model="PrivateKey" rows="6" type="textarea"
                        placeholder="PrivateKey" />

                    <div style="margin-top: 20px">
                    </div>
                    <el-input size="large" maxlength="1024" show-word-limit v-model="PublicKey" rows="6" type="textarea"
                        placeholder="PublicKey" />
                </el-card>

                <el-card shadow="hover" style="margin-top:15px">
                    <el-button type="primary" @click="GenerateRSAKeyResult">Sign</el-button>
                    <div style="margin-top: 20px">
                    </div>
                    <el-input size="large" maxlength="2048" show-word-limit rows="3" type="textarea"
                        placeholder="Sign Content" spellcheck="false" />
                    <div style="margin-top: 20px">
                    </div>

                    <el-input size="large" maxlength="3400" show-word-limit rows="6" type="textarea"
                        placeholder="Signature" />

                    <div style="margin-top: 20px">
                    </div>
                </el-card>
            </el-col>
        </div>
    </main>
</template>
  
<script lang="ts" setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus';

import { CreateRSAKeyPair } from '../service';

const PlainText = ref("");

const PrivateKey = ref("");
const PublicKey = ref("");


const GenerateRSAKeyResult = async () => {
    let resp = await CreateRSAKeyPair(4096);
    if (resp.errmsg != "") {
        ElMessage.error(resp.errmsg)
    } else {
        PrivateKey.value = resp.PrivateKey;
        PublicKey.value = resp.PublicKey;
    }
}

</script>
