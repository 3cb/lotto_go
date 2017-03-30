<template>
  <div id="app">
    <header class="header"><h1>Lotto</h1></header>
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span style="line-height: 36px;">Mega Millions</span>
      </div>
      <div>
        <mega-millions></mega-millions>
      </div>
    </el-card>
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span style="line-height: 36px;">Powerball</span>
      </div>
      <div>
      
      </div>
    </el-card>
  </div>
</template>

<script>
import MegaMillions from './components/MegaMillions.vue'

export default {
  name: 'app',
  components: {
    MegaMillions
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
  text-align: left;
  color: #2c3e50;
  margin: 60px 60px 60px 60px;
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

.header {
  width: 100%;
  padding: 0%;
  margin: 0%;
}

.box-card {
  margin: 0px 0px 10px 0px;
}

</style>
