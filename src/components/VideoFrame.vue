<template>
    <div class="container">
        <my-video :sources="video.sources" :options="video.options"></my-video>
    </div>
</template>

<script>
import myVideo from "vue-video";
export default {
  props: ["src"],
  data: function() {
    return {
      video: {
        sources: [
          {
            src: this.src,
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
