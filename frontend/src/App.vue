<template>
  <div id="app">
    <div v-if="thx">
      <NewsAPI/>
    </div>
    <searchNews v-on:news-search="handleChange"/>
    <News v-bind:news="news"/>
    <div class="loader">
      <pulse-loader :loading="loading" :color="color"></pulse-loader>
    </div>
  </div>
</template>

<script>
import Request from "request";
import News from "./components/News.vue";
import NewsAPI from "./components/NewsAPI";
import searchNews from "./components/searchNews";
import PulseLoader from "vue-spinner/src/PulseLoader.vue";

export default {
  name: "app",
  components: {
    NewsAPI,
    News,
    searchNews,
    PulseLoader
  },
  created() {
    this.getNews();
    window.addEventListener("scroll", this.handleScroll);
  },
  // updated() {},
  data() {
    return {
      thx: true,
      loading: false,
      allNews: [],
      news: [],
      limit: 0,
      color: "#FD6A02",
      search: ""
    };
  },
  methods: {
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
    getNews(search = "technology") {
      this.loading = true;
      Request(
        `http://localhost:3456/news/${search}`,
        (error, res, body) => {
          const result = JSON.parse(body);
          if (result.status === "error") {
            this.loading = false;
            alert("Not found");
          } else {
            this.allNews = result;
            this.thx = false;
            this.incLimit();
          }
        }
      );
    },
    incLimit() {
      if (this.limit < this.allNews.length) {
        this.loading = true;
        setTimeout(() => {
          this.news = [
            ...this.news,
            ...this.allNews.slice(this.limit, (this.limit += 10))
          ];
          this.loading = false;
        }, 2000);
      }
    },
    handleChange(search) {
      this.reset();
      this.getNews(search);
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
