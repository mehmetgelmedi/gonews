<template>
  <div class="search-container">
    <i id="search" class="fas fa-search fa-3x hide"></i>
    <input
      type="text"
      v-model="search"
      v-on:keyup.enter="handleSearch"
      placeholder="search news, don't forget press enter!"
      class="search-input"
    >
  </div>
</template>

<script>
export default {
  name: "searchNews",
  data() {
    return {
      search: ""
    };
  },
  methods: {
    handleSearch() {
      this.$emit("news-search", this.search);
    }
  }
};
</script>

<style scoped>
html,
body {
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
  background: #01786f;
}

.hide {
  display: none;
  padding-top: 10px;
}

.search-container {
  height: 100%;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
}

.search-input {
  width: 50%;
  height: 35px;
  font-size: 2em;
  border: 4px solid;
  padding: 10px;
  caret-color: red;
}

@media only screen and (max-width: 800px) {
  .search-input {
    width: 75%;
  }
}

.search-input:focus {
  outline-width: 0;
}

/* .search-input:hover + span > .fa-search-1 {
  animation: move-right 0.5s;
} */

.move {
  animation: move-right 1s;
}

@keyframes move-right {
  form {
    transform: translateX(0);
  }

  to {
    opacity: 0;
    transform: translateX(40px);
  }
}

.search-input::placeholder {
  opacity: 0.7;
  color: #000;
}

span {
  border: 4px solid;
  height: 35px;
  padding: 10px;
  border-left: 0;
  background: #fff;
  cursor: pointer;
}

.fa-search-1 {
  margin-top: -6.5px;
}

.fa-search-2 {
  display: block;
  position: absolute;
  animation: move 3s;
  animation-fill-mode: forwards;
}

@keyframes move {
  0% {
    opacity: 0;
    visibility: visible;
    padding-right: 50%;
  }
  10% {
    opacity: 1;
    padding-right: 30%;
  }
  30% {
    padding-right: 50%;
  }
  50% {
    padding-right: 20%;
  }
  60% {
    padding-right: 0;
  }
  80% {
    padding-right: 35%;
  }
  95% {
    opacity: 1;
  }
  100% {
    opacity: 0;
    padding-right: 55%;
    visibility: hidden;
  }
}
</style>


