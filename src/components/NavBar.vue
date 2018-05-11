<template>
    <div class="topnav">
        <ul>
          <li v-for="(item, index) in pages" :key="item">
            <a :href="item" :value="index" @click.prevent="$emit('input', index)">{{ item }}</a>
          </li>
        </ul>
    </div>
</template>

<script>
export default {
  name: "NavBar",
  props: ["pages"]
};
</script>
<style lang="less" scoped>

</style>
