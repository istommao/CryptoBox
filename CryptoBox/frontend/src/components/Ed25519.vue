
<template>
   <main>
      <div>
         Ed25519
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

</script>
