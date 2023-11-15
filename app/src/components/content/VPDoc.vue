<script setup lang="ts">

import { Content } from '@/stores/content';

let rawHtml = `
<h2 class="toc-header" id="火车">
    火车
</h2>
<ul>
    <li>可在杭州火车东站（距学校较近）、杭州站（城站）乘车，地铁直达</li>
    <li>推荐网上购票，可以在师生办事大厅取票机取票或火车站取票机取票，带好身份证</li>
    <li>
        高铁动车大部分车次<strong>全价票</strong>
        支持刷身份证直接进出站，请注意购票时列车列表身份证标识
    </li>
</ul>
<h3 class="toc-header" id="火车优惠">
    火车优惠
</h3>
<ol>
    <li>在教务系统中填好火车优惠区间，优惠区间不可随意更改</li>
    <li>优惠区间只能在每年6/1-9/30以及当年12/1-次年3/31使用共四次（单程）</li>
    <li>购买学生票需在取票时带好带优惠磁条的学生证，否则无法取票</li>
    <li>学生票只享受半价硬座客票、加快票和空调票，动车二等座打75折，软座、软卧、动车一等票均无优惠</li>
</ol>`

</script>

<template>
  <div
    class="VPDoc"
  >
    <slot name="doc-top" />
    <div class="container">

      <div class="content">
        <div class="content-container">
          <slot name="doc-before" />
            <main class="main">
              <Content class="vp-doc" :htmlContent="$route.path + rawHtml" />
            </main>
          <slot name="doc-after" />
        </div>
      </div>
    </div>
    <slot name="doc-bottom" />
  </div>
</template>

<style scoped>
.VPDoc {
  padding: 32px 24px 96px;
  width: 100%;
}

.container {
  margin: 0 auto;
  width: 100%;
}

.content {
  position: relative;
  margin: 0 auto;
  width: 100%;
  padding: 0 32px;
}

.content-container {
  margin: 0 auto;
}

.VPDoc.has-aside .content-container {
  max-width: 688px;
}

.external-link-icon-enabled :is(.vp-doc a[href*='://'], .vp-doc a[target='_blank'])::after {
  content: '';
  color: currentColor;
}

.vp-doc {
  font-size: 16px;
  line-height: 1.8;
  color: var(--vp-c-text-1);
}
</style>
