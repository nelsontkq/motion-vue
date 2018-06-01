<template>
    <div id="navigation">
        <ul>
          <li v-for="item in pages" :key="item">
            <a :href="item" :class="{active: activeItem == item}" :mouseover="hover" @click.prevent="$emit('input',activate(item))">
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
      activeItem: this.selectedPage
    };
  },
  methods: {
    activate: function(item) {
      console.log(item);
      this.activeItem = item;
      return item;
    },
    hover: function(event) {
      console.log(event);
    }
  }
};
</script>
<style lang="less" scoped>
@navigation-color: rgb(68, 68, 68);
@navigation-color-active: darken(@navigation-color, 20%);
#navigation {
  all: initial;

  overflow: hidden;
  li {
    list-style: none;
  }
  a {
    float: left;
    text-align: center;
    position: relative;
    color: #f2f2f2;
    text-decoration: none;
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
