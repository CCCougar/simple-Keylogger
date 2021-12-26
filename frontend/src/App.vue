<template>
  <div id="app">
<!--    <img alt="Keylogger logo" src="./assets/keylogger.jpg" height="10%" width="10%">-->
<!--    <HelloWorld msg="Welcome to Your Vue.js App"/>-->
<!--    <h3>一共有xxx台被种植了Keylogger的主机</h3>-->
    <show-beacons></show-beacons>
  </div>
</template>

<script>
import ShowBeacons from "@/components/ShowBeacons";

export default {
  name: 'App',
  components: {
    ShowBeacons
  },
}
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
