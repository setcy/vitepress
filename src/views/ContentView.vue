<script setup lang="ts">

import VPContent from '@/components/content/VPContent.vue';

</script>
<template>
  <div class="about">
    <VPContent></VPContent>
  </div>
</template>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>
