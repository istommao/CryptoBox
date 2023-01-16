
<template>
  <main>
    <div style="display: block; margin: 0 auto">
      <el-col >
        <el-card shadow="hover">
          <el-button type="primary" @click="GenerateKeyPair">Generate</el-button>
          <div style="margin-top: 20px">
          </div>
          <el-input size="large" maxlength="128" show-word-limit v-model="PrivateKey"  rows="3" type="textarea" placeholder="Private Key" />

          <div style="margin-top: 20px">
          </div>
          <el-input size="large"  maxlength="128" show-word-limit v-model="PubKey"  rows=3 type="textarea" placeholder="Public Key" />
        </el-card>
      </el-col>
    </div>
  </main>
</template>
  
<script lang="ts" setup>
import { ref } from 'vue'
// import { Greet } from '../../wailsjs/go/main/App'



const PrivateKey = ref("");
const PubKey = ref("");

const GenerateKeyPair = () => {
  PrivateKey.value = "asdfasdbasdf";
  PubKey.value = "asdfasdfadsf";
}



</script>
