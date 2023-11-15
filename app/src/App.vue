<script setup lang="ts">

import { RouterLink, RouterView } from 'vue-router'

</script>

<template>
  <div class="controlPanelPadding">
    <RouterView />
  </div>
</template>

<style scoped>
header {
  line-height: 1.5;
  max-height: 10vh;
}
</style>
