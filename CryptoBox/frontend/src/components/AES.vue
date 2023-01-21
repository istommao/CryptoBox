
<template>
   <main>
      <div style="display: block; margin: 0 auto">
         <el-col>
            <el-card shadow="hover">
               <el-radio-group v-model="KeySize" size="large">
                  <el-radio-button label="128" />
                  <el-radio-button label="192" />
                  <el-radio-button label="256" />
               </el-radio-group>
               <div style="margin-top: 20px">
               </div>

               <div style="display: flex">
                  <el-button size="large" @click="RandomAesKeyAction">Random AesKey</el-button>&nbsp;&nbsp;
                  <el-input size="large" maxlength="128" show-word-limit v-model="AESKey" placeholder="AES Key" />
               </div>
               <div style="margin-top: 20px">
               </div>
               <div style="display: flex">
                  <el-button size="large" @click="RandomNonceAction">Random&nbsp;&nbsp;Nonce</el-button>&nbsp;&nbsp;
                  <el-input size="large" maxlength="128" show-word-limit v-model="Nonce" placeholder="Nonce" />
               </div>
               <div style="margin-top: 20px">
               </div>
               <el-input size="large" type="textarea" rows="6" spellcheck="false" maxlength="1024" show-word-limit v-model="PlainText" placeholder="Plain Text" />
               <div style="margin: 20px 0 20px 0; display: flex">
                  <el-button size="large" type="primary" @click="EncryptResult">Encrypt</el-button>
                  &nbsp;&nbsp;
                  <el-radio-group v-model="OutputFormat" size="large">
                  <el-radio-button label="base64"  />
                  <el-radio-button label="hex" />
               </el-radio-group>
               </div>
               <el-input size="large" type="textarea" rows="6" spellcheck="false" maxlength="1024" show-word-limit v-model="CipherText" placeholder="Cipher Text" />
            </el-card>
         </el-col>
      </div>
   </main>
</template>
 
<script lang="ts" setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus';
import { Key } from '@element-plus/icons-vue';

import { AesGCMEncrypt } from '../service';

const AESKey = ref("");
const KeySize = ref(128);

const PlainText = ref("");
const CipherText = ref("");

const OutputFormat = ref("base64");

const Nonce = ref("");

function Buf2Hex(buffer: ArrayBuffer) { // buffer is an ArrayBuffer
   return [...new Uint8Array(buffer)]
      .map(x => x.toString(16).padStart(2, '0'))
      .join('');
}


const RandomNonceAction = async () => {
   const array = new Uint8Array(12);
   let resultBytes = crypto.getRandomValues(array);

   // hex format


   Nonce.value = Buf2Hex(resultBytes);
}

const RandomAesKeyAction = async () => {
   let byteSize = 16;
   if (KeySize.value == 128) {
      byteSize = 16
   } else if (KeySize.value == 192) {
      byteSize = 24
   } else if (KeySize.value == 256) {
      byteSize = 32;
   }
   const array = new Uint8Array(byteSize);
   let resultBytes = crypto.getRandomValues(array);

   // hex format


   AESKey.value = Buf2Hex(resultBytes);
}

const EncryptResult = async () => {

  let resp = await AesGCMEncrypt(AESKey.value, PlainText.value, Nonce.value, OutputFormat.value);

//   console.log(">>>>", resp);

  CipherText.value = resp.CipherText;
}

</script>
