
<template>
  <div>
    <ul>
      <li v-for="item in videos" :key="item.id">
        <a href="" :value="item.Path" @click.prevent="$emit('input', index)">{{ item.Path}}</a>
        <button :value="item" v-on:click="deleteItem(item.id)">delete</button>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "DirectoryBrowse",
  props: ["src"],
  data: function() {
    return {
      videos: []
    };
  },
  created() {
    axios.get(this.src, { crossdomain: true }).then(response => {
      // JSON responses are automatically parsed.
      this.videos = response.data;
    });
  },
  methods: {
    deleteItem: function(item) {
      axios.delete(this.src + "/" + item);
    }
  }
};
</script>
<style lang="less" scoped>

</style>
