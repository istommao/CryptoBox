
<template>
   <main>
      <div style="display: block; margin: 0 auto">
         <el-col>
            <el-card shadow="hover">
               <el-button type="primary" @click="EncryptResultAction">Encrypt</el-button>
               <div style="margin-top: 20px">
               </div>
             
               <div style="display: flex">
                  <el-button size="large" @click="RandomKeyAction">Random Key</el-button>&nbsp;&nbsp;&nbsp;&nbsp;
                  <el-input size="large" maxlength="128" show-word-limit v-model="ChaCha20Key" placeholder="Key" />
               </div>

               <div style="margin-top: 20px"></div>
               <div style="display: flex">
                  <el-button size="large" @click="RandomNonceAction">Random Nonce</el-button>&nbsp;&nbsp;
                  <el-input size="large" maxlength="128" show-word-limit v-model="ChaCha20Nonce" placeholder="Nonce" />
               </div>

               <div style="margin-top: 20px"></div>
               <el-input size="large" maxlength="128" show-word-limit v-model="Payload" rows="6" type="textarea"
                  placeholder="Payload" spellcheck="false" />

               <div style="margin-top: 20px">
               </div>
               <el-input size="large" maxlength="128" show-word-limit v-model="EncryptResult" rows="6" type="textarea"
                  placeholder="Result" />
            </el-card>
         </el-col>
      </div>
   </main>
</template>
 
<script lang="ts" setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus';

// @ts-ignore
// import * as stdlib from '@stablelib/chacha';
import JSChaCha20 from "js-chacha20";

import { CreateX25519KeyPair } from '../service';

function Buf2Hex(buffer: ArrayBuffer) { // buffer is an ArrayBuffer
   return [...new Uint8Array(buffer)]
      .map(x => x.toString(16).padStart(2, '0'))
      .join('');
}

const Payload = ref("");
const EncryptResult = ref("");

const ChaCha20Key = ref("");

const ChaCha20Nonce = ref("");


const RandomNonceAction = async () => {
   const array = new Uint8Array(12);
   let resultBytes = crypto.getRandomValues(array);


   ChaCha20Nonce.value = Buf2Hex(resultBytes);
}

const fromHexString = (hexString: any) =>
  Uint8Array.from(hexString.match(/.{1,2}/g).map((byte: string) => parseInt(byte, 16)));


const RandomKeyAction = async () => {  
   const array = new Uint8Array(32);
   let resultBytes = crypto.getRandomValues(array);

   ChaCha20Key.value = Buf2Hex(resultBytes);
}

const Uint8ToBase64String = (u8a: any) => {
  return btoa(String.fromCharCode.apply(null, u8a));
};

const EncryptResultAction = async () => {
   console.log( "====");

   // let result = stdlib.stream(fromHexString(ChaCha20Key.value),  fromHexString(ChaCha20Nonce.value), new Uint8Array() );

   let dataByte = new TextEncoder().encode(Payload.value);

   const result = new JSChaCha20(fromHexString(ChaCha20Key.value), fromHexString(ChaCha20Nonce.value)).encrypt(dataByte);

   console.log(result, "====");
   EncryptResult.value = Uint8ToBase64String(result);
}
</script>
