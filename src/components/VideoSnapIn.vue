
<template>
  <div>
    <ul>
      <li v-for="item in videos" :key="item.id">
        <input type="checkbox" :value="item" :id="item.id" v-model="checkedVideos"> 
         <a href="" @click.prevent="changeVideo(item.id)">{{ item.Name }}</a>
      </li>
    </ul>
    <button @click="deleteItems">Delete</button>
    <div class="container">
      <video width="450" controls :src="video"></video>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "VideoSnapIn",
  props: ["src"],
  data: function() {
    return {
      videos: [],
      checkedVideos: [],
      video: this.src + "/" + 1
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
    },
    changeVideo: function(id) {
      console.log(this.video + " changing video to " + this.src + "/" + id);
      this.video = this.src + "/" + id;
    }
  }
};
</script>
<style lang="less" scoped>
.container {
  max-width: 80%;
  max-height: 80%;
}
</style>

