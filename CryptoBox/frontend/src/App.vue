<template>
  <el-row class="tac">
    <el-col :span="3" style="background: white;">
      <!-- <h5 class="mb-2">Default colors</h5> -->
      <el-menu
        default-active="2"
        class="el-menu-vertical-demo"
        @open="handleOpen"
        @close="handleClose"
      >
        <el-sub-menu index="1">
          <template #title>
            <el-icon><ForkSpoon /></el-icon>
            <span>Asymmetric</span>
          </template>
          <el-menu-item-group title="X25591/Ed25519">
            <router-link to="/ed25519">
              <el-menu-item index="1-1">Ed25519</el-menu-item>
            </router-link>
            <router-link to="/x25519">
              <el-menu-item index="1-1">X25519</el-menu-item>
            </router-link>
            <el-menu-item index="1-2">KeyExchange</el-menu-item>
          </el-menu-item-group>
          <el-menu-item-group title="RSA">
            <el-menu-item index="1-3">Create KeyPair</el-menu-item>
            <el-menu-item index="1-3">Encrypt/Decrypt</el-menu-item>
          </el-menu-item-group>
         <!--  <el-sub-menu index="1-4">
            <template #title>item four</template>
            <el-menu-item index="1-4-1">item one</el-menu-item>
          </el-sub-menu> -->
        </el-sub-menu>


         <el-sub-menu index="2">
          <template #title>
            <el-icon><Watermelon /></el-icon>
            <span>Hash</span>
          </template>
            <el-menu-item-group title="SHA1/2/3">
            <el-menu-item index="1-1">SHA1</el-menu-item>
            <el-menu-item index="1-2">SHA2</el-menu-item>
            <el-menu-item index="1-2">SHA3</el-menu-item>
          </el-menu-item-group>

        </el-sub-menu>



      </el-menu>
    </el-col>
    <el-col :span="21">
        <div style="margin: 15px">
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
</script>