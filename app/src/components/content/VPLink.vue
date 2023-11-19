<script lang="ts" setup>
import {computed} from 'vue'
import {EXTERNAL_URL_RE} from '@/support/shared'
import {RouterLink} from "vue-router";

const props = defineProps<{
  tag?: string
  href: string
  noIcon?: boolean
  target?: string
  rel?: string
}>()

const tag = computed(() => props.tag ?? (props.href ? 'a' : 'span'))
const isExternal = computed(() => props.href && EXTERNAL_URL_RE.test(props.href))
</script>

<template>
  <RouterLink :to="href">
    <component
        :is="tag"
        :class="{
        link: href,
        'vp-external-link-icon': isExternal,
        'no-icon': noIcon
      }"
        :rel="rel ?? (isExternal ? 'noreferrer' : undefined)"
        :target="target ?? (isExternal ? '_blank' : undefined)"
        class="VPLink"
    >
      <slot/>
    </component>
  </RouterLink>
</template>
