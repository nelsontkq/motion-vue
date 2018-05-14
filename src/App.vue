<template>
  <div id="app">
    <NavBar :pages="pages" v-model="selectedPage"/>
    <VideoFrame v-show="selectedPage == 0" :src="src + '/' + selectedVideo"/>
    <DirectoryBrowse v-show="selectedPage == 1" :src="src" v-model="selectedVideo"/>
  </div>
</template>

<script>
import VideoFrame from "./components/VideoFrame";
import DirectoryBrowse from "./components/DirectoryBrowse";
import NavBar from "./components/NavBar";
export default {
  name: "app",
  data: function() {
    return {
      src: "http://localhost:3000/video",
      selectedVideo: 1,
      selectedPage: 0,
      pages: ["VideoFrame", "DirectoryBrowse"]
    };
  },
  components: {
    NavBar,
    VideoFrame,
    DirectoryBrowse
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
