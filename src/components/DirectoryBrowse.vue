
<template>
  <table>
    <thead>
      <tr>
        <th v-for="key in videos"
          @click="sortBy(key)"
          :class="{ active: sortKey == key }" :key="videos.Id">
          {{ key | capitalize }}
          <span class="arrow" :class="sortOrders[key] > 0 ? 'asc' : 'dsc'">
          </span>
        </th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="entry in filteredData">
        <td v-for="key in columns">
          {{entry[key]}}
        </td>
      </tr>
    </tbody>
  </table>
</template>
<script>
import axios from "axios";
export default {
  name: "DirectoryBrowse",
  props: ["src"],
  data: function() {
    return {
      videos: {}
    };
  },
  created() {
    axios.get(this.src).then(response => {
      // JSON responses are automatically parsed.
      this.videos = response.data;
    });
  }
};
</script>
<style lang="less" scoped>

</style>
