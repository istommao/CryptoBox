import { createApp } from 'vue'
import App from './App.vue'
import * as VueRouter from 'vue-router';
// import HelloWorldVue from './components/HelloWorld.vue';

import HomePage from './components/HomePage.vue';
import SHA1AndMD5Page from './components/SHA1_MD5.vue';

import SHA2Page from './components/SHA2.vue';
import SHA3Page from './components/SHA3.vue';

import Ed25519Vue from './components/Ed25519.vue';
import X25519Page from './components/X25519.vue';

import RSAPage from './components/RSA.vue';

import AESPage from './components/AES.vue';
import ChaCha20Page from './components/ChaCha20.vue';

import 'element-plus/theme-chalk/dark/css-vars.css'

import 'element-plus/dist/index.css'

const routes = [
  { path: '/', component: HomePage, name: "Home" },
  { path: '/ed25519', component: Ed25519Vue },
  { path: '/x25519', component: X25519Page },
  { path: '/rsa', component: RSAPage },

  { path: '/sha1_md5', component: SHA1AndMD5Page },
  { path: '/sha2', component: SHA2Page },
  { path: '/sha3', component: SHA3Page },

  { path: '/aes', component: AESPage },
  { path: '/chacha20', component: ChaCha20Page },
]


const router = VueRouter.createRouter({
  // 4. Provide the history implementation to use. We are using the hash history for simplicity here.
  history: VueRouter.createWebHashHistory(),
  routes, // short for `routes: routes`
})


const app = createApp(App)


app.use(router)

app.mount('#app')
