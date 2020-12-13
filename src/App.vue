<template>
  <div id="app">
    <img alt="Kubernetes logo" src="./assets/kubernetes.png" />
    <Hello msg="Halo! Selamat Datang di KuDemo App" />
  </div>
</template>

<script>
import Hello from "./components/Hello.vue";

export default {
  name: "App",
  components: {
    Hello,
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
