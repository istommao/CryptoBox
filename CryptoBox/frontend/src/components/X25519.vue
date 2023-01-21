
<template>
   <main>
      <div style="display: block; margin: 0 auto">
         <el-col>
            <el-card shadow="hover">
               <el-button type="primary" @click="GenerateKeyResult">Generate</el-button>
               <div style="margin-top: 20px">
               </div>
               <el-input size="large" maxlength="128" show-word-limit v-model="PrivateKey" rows="6" type="textarea"
                  placeholder="PrivateKey" />

               <div style="margin-top: 20px">
               </div>
               <el-input size="large" maxlength="128" show-word-limit v-model="PublicKey" rows="6" type="textarea"
                  placeholder="PublicKey" />
            </el-card>
         </el-col>
      </div>
   </main>
</template>
 
<script lang="ts" setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus';

import { CreateX25519KeyPair } from '../service';


const PrivateKey = ref("");
const PublicKey = ref("");


const GenerateKeyResult = async () => {
   let resp = await CreateX25519KeyPair();

   PrivateKey.value = resp.PrivateKey;
   PublicKey.value = resp.PublicKey;

}

</script>
