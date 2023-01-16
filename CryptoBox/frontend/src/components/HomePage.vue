
<template>
    <main>
       <div>
          Hello CryptoBox

          <!-- <a href="#" @click="OpenBrowserURL" title="Hello密码学">Hello 密码学</a> -->
       </div>
    </main>
 </template>
   
 <script lang="ts" setup>
 import { reactive } from 'vue'
 import { Greet } from '../../wailsjs/go/main/App'
 
 const data = reactive({
    name: "",
    resultText: "Please enter your name below 👇",
 })
 
 function greet() {
    Greet(data.name).then(result => {
       data.resultText = result
    })
 }

//  const OpenBrowserURL = () => {
//     let url = "https://hello-crypto.cn/"
//     BrowserOpenURL(url)
//  }
 
 </script>
 