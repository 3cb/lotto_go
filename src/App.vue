<template>
  <div id="app">
    <el-card class="box-card">
      <el-tabs v-model="activeName">
        <el-tab-pane label="MegaMillions" name="first">MegaMillions</el-tab-pane>
        <el-tab-pane label="Powerball" name="second">Powerball</el-tab-pane>
      </el-tabs>
    </el-card>
  </div>
</template>

<script>
export default {
  name: 'app',
  components:{
    // MegaMillions
  },
  data () {
    return {
      activeName: 'first'
    }
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

h1, h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}
</style>
