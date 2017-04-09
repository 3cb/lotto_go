<template>
    <div>
        <el-col :span="9">
            <div class="block grid-content">
                <span class="bold">Set Number of Picks:</span>
                <br>
                <br>
                <el-slider v-model="count" :min="1">
                </el-slider>
                <br>
                <el-button type="primary" @click="generatePicks">Generate Picks
                </el-button>
            </div>
        </el-col>
        <el-col :span="15">
            <div class="block">
                <el-table
                    class="pickTable"
                    :data="picks"
                    height="150"
                    style="width: 80%">
                    <el-table-column
                        prop="numbers"
                        label="Numbers"
                        align="center">
                    </el-table-column>
                    <el-table-column
                        prop="ball"
                        :label="ballLabel"
                        align="center">
                    </el-table-column>
                </el-table>
            </div>
        </el-col>
    </div>
</template>

<script>
export default {
    props: ['ballLabel'],
    data() {
        return {
            count: 1,
            picks: []
        }
    },
    methods: {
        generatePicks() {
            var mult
            var multb
            var nums = []

            if (this.ballLabel === 'Mega Ball') {
                mult = 75
                multb = 15
            } else {
                mult = 69
                multb = 26
            }

            for (var j = 0; j < this.count; j++) {
                let pick = {
                    numbers: [],
                    ball: ''
                }
                let intArr = []
                let strArr = []

                for (var i = 0; i < 5; i++) {
                    intArr.push(Math.floor((Math.random() * 75) + 1))
                }
                strArr = intArr.map(x => {
                    return x.toString()
                })
                pick.numbers = strArr.join(' ')
                let x = Math.floor((Math.random() * 15) + 1)
                x.toString()
                pick.ball = x
                nums.push(pick)
            }
            this.picks = nums
        }
    }
}
</script>

<style>
    .pickTable {
        float: right;
    }
    .bold {
        font-weight: bold;
    }
</style>