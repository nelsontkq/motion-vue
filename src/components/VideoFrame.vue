<template>
    <div class="container">
        <my-video :sources="video.sources" :options="video.options"></my-video>
    </div>
</template>

<script>
import myVideo from "vue-video";
export default {
  data: function() {
    return {
      video: {
        sources: [
          {
            src: "http://localhost:8081/video/1",
            type: "video/mp4"
          }
        ],
        options: {
          volume: 0.6
        }
      }
    };
  },
  components: {
    myVideo
  }
};
</script>

<style lang="less" scoped>
.container {
  max-width: 80%;
  max-height: 80%;
}
</style>
