<template>
  <div>
    <div class="row" v-bind:key="new Date().getTime()+Date.parse(n.publishedAt)" v-for="n in news">
      <NewsItem v-bind:n="n"/>
    </div>
  </div>
</template>

<script>
import NewsItem from "./NewsItem.vue";

export default {
  name: "News",
  props: ["news"],
  components: {
    NewsItem
  }
};
</script>

<style lang="scss" scoped>
@import url(https://fonts.googleapis.com/css?family=Open+Sans:300,400,700);
// Variables
$regal-blue: #034378;
$san-juan: #2d4e68;
$open-sans: "Open Sans", sans-serif;

* {
  box-sizing: border-box;
}

body {
  background-image: linear-gradient(to right, $regal-blue, $san-juan);
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  height: 100vh;
  font-family: $open-sans;
}

a {
  text-decoration: none;
}

h1 {
  font-family: $open-sans;
  font-weight: 300;
}

.row {
  max-width: 900px;
  margin: 50px auto 0;
}
</style>

