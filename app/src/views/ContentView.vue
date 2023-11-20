<script lang="ts" setup>

import VPSidebar from "@/components/content/VPSidebar.vue";
import VPNavBar from "@/components/content/VPNavBar.vue";
import {useSidebarControl} from "@/stores/sidebar";
import {watch} from "vue";
import {useRoute} from "vue-router";
import VPBackdrop from "@/components/content/VPBackdrop.vue";

const {
  isOpen: isSidebarOpen,
  open: openSidebar,
  close: closeSidebar
} = useSidebarControl()

const route = useRoute()
watch(() => route.path, closeSidebar)

</script>
<template>
  <VPBackdrop :show="isSidebarOpen" class="backdrop" @click="closeSidebar"/>
  <VPNavBar :open="isSidebarOpen" @open-menu="openSidebar"></VPNavBar>
  <VPSidebar :open="isSidebarOpen" has-sidebar></VPSidebar>
  <RouterView/>
</template>

<style>

</style>
