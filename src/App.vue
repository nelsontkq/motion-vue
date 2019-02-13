<template>
  <div id="app">
      <top-nav :pages="pages" v-model="selectedPage"/>
    <live-stream v-show="selectedPage == pages[0]" :streamSrc="streamSrc" />
    <video-snap v-show="selectedPage == pages[1]" :src="src"/>
  </div>
</template>

<script>
import VideoSnapIn from "./components/VideoSnapIn";
import NavBar from "./components/NavBar";
import LiveStream from "./components/LiveStream";
export default {
  name: "app",
  data: function() {
    return {
      src: "http://localhost:5000/videos",
      streamSrc: "http://localhost:8081",
      selectedPage: "LiveStream",
      pages: ["LiveStream", "VideoSnapIn"]
    };
  },
  components: {
    "top-nav": NavBar,
    "video-snap": VideoSnapIn,
    "live-stream": LiveStream
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
