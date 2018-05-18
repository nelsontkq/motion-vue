<template>
    <div id="navigation">
        <ul>
          <li v-for="(item, index) in pages" :key="item">
            <a :href="item" :class="{active: activeIndex == index}" :mouseover="hover" :value="index" @click.prevent="$emit('input',activate(index))">
              {{ item }}
            </a>
          </li>
        </ul>
    </div>
</template>

<script>
export default {
  name: "NavBar",
  props: ["pages", "selectedPage"],
  data: function() {
    return {
      activeIndex: 0
    };
  },
  methods: {
    activate: function(sender) {
      console.log($event);
      this.activeIndex = index;
      return index;
    },
    hover: function(event) {}
  }
};
</script>
<style lang="less" scoped>
@navigation-color: rgb(68, 68, 68);
@navigation-color-active: darken(@navigation-color, 20%);
#navigation {
  all: initial;
  display: block;
  overflow: hidden;

  a {
    color: #f2f2f2;
    text-decoration: none;
    text-align: center;
    cursor: default;
    padding: 12px 14px;
    background-color: @navigation-color;
  }
  a:hover {
    background-color: @navigation-color-active;
  }
  a.active {
    background-color: @navigation-color-active;
  }
}
</style>
