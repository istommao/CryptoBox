
<template>
    <main>
        <div style="display: block; margin: 0 auto">
            <el-col>
                <el-card shadow="hover">
                    <el-button type="primary" @click="GenerateSHAResult">Generate Hash</el-button>
                    <div style="margin-top: 20px">
                    </div>
                    <el-input size="large" maxlength="128" show-word-limit v-model="PlainText" rows="3" type="textarea"
                        placeholder="Plain Text"  spellcheck="false" />
                    <div style="margin-top: 20px">
                    </div>

                    <el-input size="large" maxlength="128" show-word-limit v-model="SHA3_224" rows="3" type="textarea"
                        placeholder="SHA3 224" />

                    <div style="margin-top: 20px">
                    </div>
                    <el-input size="large" maxlength="128" show-word-limit v-model="SHA3_256" rows="3" type="textarea"
                        placeholder="SHA3 256" />

                    <div style="margin-top: 20px">
                    </div>
                    <el-input size="large" maxlength="128" show-word-limit v-model="SHA3_384" rows="3" type="textarea"
                        placeholder="SHA3 384" />

                    <div style="margin-top: 20px">
                    </div>
                    <el-input size="large" maxlength="128" show-word-limit v-model="SHA3_512" rows=3 type="textarea"
                        placeholder="SHA3 512" />
                </el-card>
            </el-col>
        </div>
    </main>
</template>
  
<script lang="ts" setup>
import { ref } from 'vue'
import { CreateSHA3Hash } from '../../wailsjs/go/main/App';

const PlainText = ref("");

const SHA3_224 = ref("");
const SHA3_256 = ref("");
const SHA3_384 = ref("");
const SHA3_512 = ref("");

const GenerateSHAResult = async() => {
    let resp = await CreateSHA3Hash(PlainText.value);
    SHA3_224.value = resp.sha3_224;
    SHA3_256.value = resp.sha3_256;
    SHA3_384.value = resp.sha3_384;
    SHA3_512.value = resp.sha3_512;
}

</script>
