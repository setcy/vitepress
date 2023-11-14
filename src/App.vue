<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import ControlPanel from './components/ControlPanel.vue'
</script>

<template>
  <header>
    <ControlPanel />
  </header>
  <RouterView />
</template>

<style scoped>
header {
  line-height: 1.5;
  max-height: 10vh;
}

.logo {
  display: block;
  width: 20px;
  margin: 0 auto 2rem;
}
</style>
