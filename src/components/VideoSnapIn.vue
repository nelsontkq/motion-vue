
<template>
<div>
<div id="frame" @keyup.delete="warned ? deleteItems : modalShow = !modalShow" >
  <div v-show="selecting" id="list" >
    <ul>
      <li v-for="item in videos" :key="item.id">
        <input type="checkbox" :value="item" :id="item.id" v-model="selectedVideos"> 
         <a href="" @click.prevent="changeVideo(item.id)">{{ item.name }}</a>
      </li>
    </ul>
  </div>
  <div v-show="!selecting" id="container">
      <button @click="selecting = true">browse...</button>
      <video width="450" controls :src="video"></video>
  </div>
</div>
<!-- Modal Component-->
  <b-modal id="warningModal" v-model="modalShow" title="Warning">
    <p class="my-4">Are you sure you want to delete these files?</p>
    <div slot="modal-footer" class="w-100">
      <b-btn size="sm" class="float-right mt-3" variant="primary" @click="modalShow=false">Cancel</b-btn>
      <b-btn size="sm" class="float-right mt-3" variant="primary" @click="modalShow=false;deleteItems()">Yes</b-btn>
    </div>
  </b-modal>
</div>
</template>

<script>
import axios from "axios";
import bModal from "bootstrap-vue/es/components/modal/modal";
import bModalDirective from "bootstrap-vue/es/directives/modal/modal";
import "bootstrap/dist/css/bootstrap.css";
import "bootstrap-vue/dist/bootstrap-vue.css";
export default {
  name: "VideoSnapIn",
  props: ["src"],
  components: {
    "b-modal": bModal
  },
  directives: {
    "b-modal": bModalDirective
  },
  data: function() {
    return {
      videos: [],
      selectedVideos: [],
      video: this.src + "/" + 1,
      selecting: true,
      modalShow: false
    };
  },
  created() {
    axios.get(this.src, { crossdomain: true }).then(response => {
      this.videos = response.data;
    });
  },
  methods: {
    deleteItems: function() {
      for (var i = 0; i < this.selectedVideos.length; i++) {
        axios.delete(this.src + "/" + this.selectedVideos[i].id, {
          crossdomain: true
        });
        var index = this.videos.indexOf(this.selectedVideos[i]);
        this.videos.splice(index, 1);
        this.selectedVideos.splice(i, 1);
      }
    },
    changeVideo: function(id) {
      this.selecting = false;
      this.video = this.src + "/" + id;
    }
  }
};
</script>
<style lang="less" scoped>
#frame {
  position: absolute;
  left: 0px;
  width: 100%;
  top: 105px;
  background: black;
  #list {
    display: block;
    float: left;
    margin-left: 20px;
    a {
      text-decoration: none;
      color: white;
      font-size: 1.2em;
    }
    li {
      overflow: hidden;
      list-style: none;
    }
  }
  #container {
    top: 0px;
    max-width: 80%;
    max-height: 80%;
    display: inline-flex;
  }
}
</style>

