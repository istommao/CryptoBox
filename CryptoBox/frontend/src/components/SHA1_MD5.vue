
<template>
    <main>
        <div style="display: block; margin: 0 auto">
            <el-col>
                <el-card shadow="hover">
                    <el-button type="primary" @click="GenerateHashResult">Generate Hash</el-button>
                    <div style="margin-top: 20px">
                    </div>
                    <el-input size="large" maxlength="128" show-word-limit v-model="PlainText" rows="3" type="textarea"
                        placeholder="Plain Text" spellcheck="false" />
                    <div style="margin-top: 20px">
                    </div>

                    <el-input size="large" maxlength="128" show-word-limit v-model="MD5" rows="3" type="textarea"
                        placeholder="MD5" />

                    <div style="margin-top: 20px">
                    </div>
                    <el-input size="large" maxlength="128" show-word-limit v-model="SHA1" rows="3" type="textarea"
                        placeholder="SHA1" />
                </el-card>
            </el-col>
        </div>
    </main>
</template>
  
<script lang="ts" setup>
import { ref } from 'vue'
import { CreateSHA1MD5Hash } from '../service';

const PlainText = ref("");

const MD5 = ref("");
const SHA1 = ref("");


const GenerateHashResult = async () => {
    let resp = await CreateSHA1MD5Hash(PlainText.value);
    MD5.value = resp.md5;
    SHA1.value = resp.sha1;

}

</script>
