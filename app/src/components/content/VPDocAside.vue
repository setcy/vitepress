<script setup lang="ts">
import VPDocAsideOutline from './VPDocAsideOutline.vue'
import type {Aside} from "@/stores/content";

const props = defineProps<{
  data: Aside[]
}>()

</script>

<template>
  <div class="VPDocAside">
    <slot name="aside-top" />

    <slot name="aside-outline-before" />
      <VPDocAsideOutline :data="data" />
    <slot name="aside-outline-after" />

    <div class="spacer" />

    <slot name="aside-bottom" />
  </div>
</template>

<style scoped>
.VPDocAside {
  display: flex;
  flex-direction: column;
  flex-grow: 1;
}

.spacer {
  flex-grow: 1;
}

.VPDocAside :deep(.spacer + .VPDocAsideSponsors),
.VPDocAside :deep(.spacer + .VPDocAsideCarbonAds) {
  margin-top: 24px;
}

.VPDocAside :deep(.VPDocAsideSponsors + .VPDocAsideCarbonAds) {
  margin-top: 16px;
}
</style>
