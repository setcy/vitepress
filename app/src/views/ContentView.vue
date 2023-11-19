<script lang="ts" setup>

import VPSidebar from "@/components/content/VPSidebar.vue";
import VPNavBar from "@/components/content/VPNavBar.vue";
</script>
<template>
  <VPNavBar></VPNavBar>
  <VPSidebar has-sidebar open></VPSidebar>
  <RouterView/>
</template>

<style>

</style>
