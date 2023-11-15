<script setup lang="ts">

import VPSidebar from "@/components/content/VPSidebar.vue";
import VPNavBar from "@/components/content/VPNavBar.vue";

</script>
<template>
  <VPNavBar></VPNavBar>
  <VPSidebar open has-sidebar></VPSidebar>
  <RouterView />
</template>

<style>

</style>
