
<template>
  <div>
    <ul>
      <li v-for="item in videos" :key="item.id">
        <input type="checkbox" :value="item" :id="item.id" v-model="checkedVideos"> 
         {{item.Path}}
      </li>
    </ul>
    <button @click="deleteItems">Delete</button>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "DirectoryBrowse",
  props: ["src"],
  data: function() {
    return {
      videos: [],
      checkedVideos: []
    };
  },
  created() {
    axios.get(this.src, { crossdomain: true }).then(response => {
      // JSON responses are automatically parsed.
      this.videos = response.data;
    });
  },
  methods: {
    deleteItems: function() {
      for (var i = 0; i < this.checkedVideos.length; i++) {
        axios.delete(this.src + "/" + this.checkedVideos[i].id, {
          crossdomain: true
        });
        var index = this.videos.indexOf(this.checkedVideos[i]);
        this.videos.splice(index, 1);
        this.checkedVideos.splice(i, 1);
      }
    }
  }
};
</script>
<style lang="less" scoped>
</style>
