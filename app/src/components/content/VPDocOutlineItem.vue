<script lang="ts" setup>
import type {Aside} from "@/stores/content";
import {easeInOutCubic} from "@/support/shared";

defineProps<{
  headers: Aside[]
  root?: boolean
}>()

function onClick(event: Event) {
  event.preventDefault();

  const el = event.target as HTMLAnchorElement;
  const id = el.href.split('#')[1];
  const section = document.getElementById(decodeURIComponent(id));

  if (section) {
    const startPosition = window.scrollY;

    const endPosition = section.offsetTop;

    const distance = endPosition - startPosition;
    const duration = 500; // Scroll duration in milliseconds
    let start: number;

    function step(timestamp: number) {
      if (!start) start = timestamp;
      const progress = timestamp - start;
      const pace = easeInOutCubic(progress / duration);
      window.scrollTo(0, startPosition + distance * pace);

      if (progress < duration) window.requestAnimationFrame(step);
    }

    window.requestAnimationFrame(step);
  }
}

</script>

<template>
  <ul :class="root ? 'root' : 'nested'">
    <li v-for="{ children, anchor, title } in headers">
      <a :href="anchor" :title="title" class="outline-link" @click="onClick">{{ title }}</a>
      <template v-if="children?.length">
        <VPDocOutlineItem :headers="children"/>
      </template>
    </li>
  </ul>
</template>

<style scoped>
.root {
  position: relative;
  z-index: 1;
}

.nested {
  padding-left: 16px;
}

.outline-link {
  display: block;
  line-height: 28px;
  color: var(--vp-c-text-2);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  transition: color 0.5s;
  font-weight: 400;
}

.outline-link:hover,
.outline-link.active {
  color: var(--vp-c-text-1);
  transition: color 0.25s;
}

.outline-link.nested {
  padding-left: 13px;
}
</style>
