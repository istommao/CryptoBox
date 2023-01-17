
<template>
 <main>
    <div style="display: block; margin: 0 auto">
      <el-col>
        <el-card shadow="hover">
          <el-button type="primary" @click="GenerateSHAResult">Generate Hash</el-button>
          <div style="margin-top: 20px">
          </div>
          <el-input size="large" maxlength="128" show-word-limit v-model="PlainText" rows="3" type="textarea"
            placeholder="Plain Text" />
            <div style="margin-top: 20px">
          </div>

          <el-input size="large" maxlength="128" show-word-limit v-model="SHA224" rows="3" type="textarea"
            placeholder="SHA 224" />

          <div style="margin-top: 20px">
          </div>
          <el-input size="large" maxlength="128" show-word-limit v-model="SHA256" rows="3" type="textarea"
            placeholder="SHA 256" />

          <div style="margin-top: 20px">
          </div>
          <el-input size="large" maxlength="128" show-word-limit v-model="SHA384" rows="3" type="textarea"
            placeholder="SHA 384" />

          <div style="margin-top: 20px">
          </div>
          <el-input size="large" maxlength="128" show-word-limit v-model="SHA512" rows=3 type="textarea"
            placeholder="SHA 512" />
        </el-card>
      </el-col>
    </div>
  </main>
 </template>
   
 <script lang="ts" setup>
 import { ref } from 'vue'
import { CreateSHA2Hash } from '../../wailsjs/go/main/App';

 const PlainText = ref("");

 const SHA224 = ref("");
 const SHA256 = ref("");
 const SHA384 = ref("");
 const SHA512 = ref("");

 const GenerateSHAResult = async() => {
    let resp = await CreateSHA2Hash(PlainText.value);
    SHA224.value = resp.sha224;
    SHA256.value = resp.sha256;
    SHA384.value = resp.sha384;
    SHA512.value = resp.sha512;
}

 </script>
 