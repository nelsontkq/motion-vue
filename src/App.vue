<template>
  <div id="app">
    <NavBar :pages="pages" v-model="selectedPage"/>
    <LiveStream v-show="selectedPage == 0" :streamSrc="streamSrc" />
    <DirectoryBrowse v-show="selectedPage == 1" :src="src"/>
  </div>
</template>

<script>
import DirectoryBrowse from "./components/VideoSnapIn";
import NavBar from "./components/NavBar";
export default {
  name: "app",
  data: function() {
    return {
      src: "http://localhost:3000/video",
      streamSrc: "http://localhost:8081",
      selectedPage: 0,
      pages: ["LiveStream", "VideoSnapIn"]
    };
  },
  components: {
    NavBar,
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
