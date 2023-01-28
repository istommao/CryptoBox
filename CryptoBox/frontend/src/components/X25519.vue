
<template>
   <main>
      <div style="display: block; margin: 0 auto">
         <el-col>
            <el-card shadow="hover">
               <el-button type="primary" @click="GenerateKeyResult">Generate</el-button>
               <div style="margin-top: 20px">
               </div>
               <span class="ml-3 w-35 text-gray-600 inline-flex items-center">Alice PrivateKey</span>
               <el-input size="large" maxlength="128" show-word-limit v-model="AlicePrivateKey" placeholder="PrivateKey" />

               <div style="margin-top: 20px">
               </div>
               <span class="ml-3 w-35 text-gray-600 inline-flex items-center">Alice PublicKey</span>
               <el-input size="large" maxlength="128" show-word-limit v-model="AlicePublicKey" placeholder="PublicKey" />

               <div style="margin-top: 20px">
               </div>

               <span class="ml-3 w-35 text-gray-600 inline-flex items-center">Bob PrivateKey</span>
               <el-input size="large" maxlength="128" show-word-limit v-model="BobPrivateKey" placeholder="PrivateKey" />

               <div style="margin-top: 20px">
               </div>
               <span class="ml-3 w-35 text-gray-600 inline-flex items-center">Bob PublicKey</span>
               <el-input size="large" maxlength="128" show-word-limit v-model="BobPublicKey" placeholder="PublicKey" />
            </el-card>
            <div style="margin-top: 20px">
               </div>


            <el-card shadow="hover">
               <el-button type="primary" @click="GenerateShareKeyResult">Get ShareKey</el-button>
               <div style="margin-top: 20px">
               </div>
               <el-input size="large" maxlength="128" show-word-limit v-model="ShareKey" placeholder="ShareKey" />
            </el-card>
         </el-col>
      </div>
   </main>
</template>
 
<script lang="ts" setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus';

import { CreateX25519KeyPair } from '../service';

// Common.js and ECMAScript Modules (ESM)
import * as ed from '@noble/ed25519';

const Uint8ToBase64String = (u8a: any) => {
  return btoa(String.fromCharCode.apply(null, u8a));
};

const ByteArrayToHexString = (byteArray: Uint8Array) => {
  return Array.from(byteArray, function (byte) {
    return ('0' + (byte & 0xff).toString(16)).slice(-2);
  }).join('');
};


const GetShareKeyResult = async (privateKeyA: any, pubKeyB: any, format: string) => {
  try {
    const shared = await ed.curve25519.scalarMult(privateKeyA, pubKeyB);
    if (format === 'hex') {
      return ByteArrayToHexString(shared);
    } else {
      return Uint8ToBase64String(shared);
    }
  } catch (e: any) {
    return e.toString();
  }
};

const ShareKey = ref("");

const AlicePrivateKey = ref("");
const AlicePublicKey = ref("");


const BobPrivateKey = ref("");
const BobPublicKey = ref("");

const GenerateKeyResult = async () => {
   let resp = await CreateX25519KeyPair();

   AlicePrivateKey.value = resp.PrivateKey;
   AlicePublicKey.value = resp.PublicKey;

   resp = await CreateX25519KeyPair();

   BobPrivateKey.value = resp.PrivateKey;
   BobPublicKey.value = resp.PublicKey;
}

const GenerateShareKeyResult = async() => {
   let result = await GetShareKeyResult(AlicePrivateKey.value, BobPublicKey.value, "hex");
   ShareKey.value = result;
}
</script>
