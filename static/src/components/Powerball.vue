<template>
    <el-row>
    <el-collapse accordion>
        <el-collapse-item title="Search Winners" name="1">
            <search-box ballLabel="Powerball" multiplier="PowerPlay" ballProp="powerball" :api="api"></search-box>
            <br>
        </el-collapse-item>
        <el-collapse-item title="Generate Picks" name="2">
            <pick-box ballLabel="Powerball"></pick-box>
        </el-collapse-item>
        <el-collapse-item title="Lottery Info" name="3">
            <info-box
              :howLink="howLink"
              :whereLink="whereLink"
              :faq="faq">
            </info-box>
        </el-collapse-item>
    </el-collapse>
  </el-row>
</template>

<script>
import SearchBox from './SearchBox.vue'
import PickBox from './PickBox.vue'
import InfoBox from './InfoBox.vue'

export default {
    components: {
        SearchBox,
        PickBox,
        InfoBox
    },
    data() {
        return {
            api: 'https://45.79.150.247:3000/api/powerball',
            howLink: 'http://www.powerball.com/powerball/pb_howtoplay.asp',
            whereLink: 'http://www.powerball.com/powerball/pb_map.asp',
            faq: 'http://www.powerball.com/pb_contact.asp'
        }
    }
}
</script>

<style lang="css">

</style>