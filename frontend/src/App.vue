<template>
  <div id="app">
    <div v-if="this.thxNewsApi">
      <NewsAPI/>
    </div>
    <searchNews v-on:news-search="handleChange"/>
    <News v-bind:news="news"/>
    <div class="loader">
      <pulse-loader :loading="this.isSpinloading" :color="color"></pulse-loader>
    </div>
  </div>
</template>

<script>
import News from "./components/News.vue";
import NewsAPI from "./components/NewsAPI";
import searchNews from "./components/searchNews";
import PulseLoader from "vue-spinner/src/PulseLoader.vue";
import { mapState, mapMutations, mapActions } from "vuex";

export default {
  name: "app",
  components: {
    NewsAPI,
    News,
    searchNews,
    PulseLoader
  },
  computed: mapState([
    "thxNewsApi",
    "isSpinloading",
    "allNews",
    "news",
    "limit",
    "search"
  ]),
  created() {
    this.getNews();
    window.addEventListener("scroll", this.handleScroll);
  },
  // updated() {},
  data() {
    return {
      color: "#FD6A02"
    };
  },
  methods: {
        ...mapActions(["getNews", "incLimit"]),
    ...mapMutations(["spinLoader", "thankYouNewsApi"]),
    handleScroll() {
      const totalHeight = document.documentElement.scrollHeight;
      const clientHeight = document.documentElement.clientHeight;
      const scrollTop =
        document.body && document.body.scrollTop
          ? document.body.scrollTop
          : document.documentElement.scrollTop;
      if (totalHeight == scrollTop + clientHeight) {
        this.incLimit();
      }
    },
    handleChange() {
      this.reset();
      this.getNews();
    },
    reset() {
      this.news = this.allNews = [];
      this.limit = 0;
    }
  },
  destroyed() {
    window.removeEventListener("scroll", this.handleScroll);
    this.news = [];
  }
};
</script>

<style>
* {
  padding: 0;
  margin: 0;
}

body {
  font-family: Arial, Helvetica, sans-serif;
  background-color: #111111;
}
.loader {
  margin: 0 auto;
  width: 100px;
  display: block;
}
</style>
