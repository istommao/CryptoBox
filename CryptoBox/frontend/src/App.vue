<template>
  <el-row class="tac" style="border-top: 1px solid #ececec">
    <el-col :span="4" style="border-right: 1px solid var(--el-menu-border-color)">
      <!-- <h5 class="mb-2">Default colors</h5> -->
      <el-menu default-active="1" :default-openeds="DefaultOpens" class="el-menu-vertical-demo" @open="handleOpen"
        @close="handleClose">
        <el-sub-menu index="1">
          <template #title>
            <el-icon>
              <ForkSpoon />
            </el-icon>
            <span>Asymmetric</span>
          </template>
          <el-menu-item-group title="X25591/Ed25519">
            <router-link to="/ed25519">
              <el-menu-item index="1-1">Ed25519</el-menu-item>
            </router-link>
            <router-link to="/x25519">
              <el-menu-item index="1-2">X25519</el-menu-item>
            </router-link>
            <!-- <el-menu-item index="1-3">KeyExchange</el-menu-item> -->
          </el-menu-item-group>
          <el-menu-item-group title="RSA">
            <router-link to="/rsa">
              <el-menu-item index="1-4">Sign/Verify</el-menu-item>
            </router-link>

            <el-menu-item index="1-5">Encrypt/Decrypt</el-menu-item>
          </el-menu-item-group>
          <!--  <el-sub-menu index="1-4">
            <template #title>item four</template>
            <el-menu-item index="1-4-1">item one</el-menu-item>
          </el-sub-menu> -->
        </el-sub-menu>

        <el-sub-menu index="2">
          <template #title>
            <el-icon>
              <Watermelon />
            </el-icon>
            <span>Symmetric</span>
          </template>
          <el-menu-item-group title="AES">
            <router-link to="/aes">
              <el-menu-item index="2-1">AES</el-menu-item>
            </router-link>
          </el-menu-item-group>
          <el-menu-item-group title="ChaCha20/XChaCha20">
            <router-link to="/chacha20">
              <el-menu-item index="2-2">ChaCha20</el-menu-item>
            </router-link>
            <router-link to="/chacha20">
              <el-menu-item index="2-3">XChaCha20</el-menu-item>
            </router-link>
          </el-menu-item-group>

          <el-menu-item-group title="Salsa20/XSalsa20">
            <router-link to="/chacha20">
              <el-menu-item index="2-4">Salsa20</el-menu-item>
            </router-link>
            <router-link to="/chacha20">
              <el-menu-item index="2-5">XSalsa20</el-menu-item>
            </router-link>
          </el-menu-item-group>
        </el-sub-menu>

        <el-sub-menu index="3">
          <template #title>
            <el-icon>
              <Watermelon />
            </el-icon>
            <span>Hash</span>
          </template>
          <el-menu-item-group title="SHA1/2/3">
            <router-link to="/sha1_md5">
              <el-menu-item index="3-1">SHA1/MD5</el-menu-item>
            </router-link>

            <router-link to="/sha2">
              <el-menu-item index="3-2">SHA2</el-menu-item>
            </router-link>
            <router-link to="/sha3">
              <el-menu-item index="3-3">SHA3</el-menu-item>
            </router-link>
            <router-link to="/sha3">
              <el-menu-item index="3-4">RIPEMD</el-menu-item>
            </router-link>
          </el-menu-item-group>
        </el-sub-menu>
      </el-menu>

    </el-col>
    <el-col :span="20">
      <div class="MainBox">
        <router-view></router-view>
      </div>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
import {
  Document,
  ForkSpoon,
  Menu as IconMenu,
  Watermelon,
  Setting,
} from '@element-plus/icons-vue'
const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}

const DefaultOpens = ["1", "2"]
</script>