<template>
    <div>
        <el-col :span="9">
            <div class="block grid-content">
                <el-select v-model="searchRange" placeholder="Select Search Type">
                    <el-option
                        v-for="item in options"
                        :label="item.label"
                        :value="item.value">
                    </el-option>
                </el-select>
                <br>
                <el-date-picker
                    v-model="value1"
                    type="date"
                    placeholder="Pick a Day"
                    v-show="showSingleCalendar">
                </el-date-picker>
                <el-date-picker
                    v-model="value1"
                    type="daterange"
                    placeholder="Pick a Range"
                    v-show="showRangeCalendar">
                </el-date-picker>
                <br>
                <el-button type="primary" icon="search" @click="validateSingleSearch" v-show="showSingleCalendar">Search</el-button>
                <el-button type="primary" icon="search" @click="searchBetween" v-show="showRangeCalendar">Search</el-button>
            </div>
        </el-col>
        <el-col :span="15">
          <div class="block">
            <el-table
                :data="tableData"
                height="150"
                style="width: 100%">
                <el-table-column
                    prop="draw_date"
                    label="Date"
                    align="center">
                </el-table-column>
                <el-table-column
                    prop="winning_numbers"
                    label="Numbers"
                    align="center">
                </el-table-column>
                <el-table-column
                    :prop="ballProp"
                    :label="ballLabel"
                    align="center">
                </el-table-column>
                <el-table-column
                    prop="multiplier"
                    :label="multiplier"
                    align="center">
                </el-table-column>
            </el-table>
          </div>
        </el-col>
        <br>
        <el-alert
            title="Error Alert"
            type="error"
            :closable="false"
            :description="alertMessage"
            show-icon
            v-show="alertShow">
        </el-alert>
    </div>
</template>

<script>
export default {
    props: ['api', 'ballLabel', 'multiplier', 'ballProp'],
    data() {
        return {
            options: [{
                label: 'Single Day',
                value: 'Single Day'
            },
            {
                label: 'Range of Days',
                value: 'Range of Days'
            }],
            searchRange: 'Single Day',
            value1: '',
            tableData: [],
            errorMsg: '',
            alertShow: false
        }
    },
    computed: {
        showSingleCalendar: function() {
            if(this.searchRange === 'Single Day') {
                return true
            } else {
                return false
            }
        },
        showRangeCalendar: function() {
            if(this.searchRange === 'Range of Days') {
                return true
            } else {
                return false
            }
        },
        alertMessage: function() {
            if (this.ballLabel === 'Mega Ball') {
                return 'Mega Millions is drawn on Tuesdays and Fridays. Please choose another day.'
            } else {
                return 'Powerball is drawn on Wednesdays and Saturdays. Please choose another day.'
            }
        }
    },
    methods: {
        searchSingle() {
            var query = []
            query.push(this.value1)
            this.axios.post(this.api, {
                dates: query
            }).then(response => {
                this.tableData = response.data
            }).catch(error => {
                this.errorMsg = 'Error retrieving search results.'
                this.tableData = []
            })
        },
        searchBetween() {
            var query = this.value1
            this.axios.post(this.api, {
                dates: query
            }).then(response => {
                this.tableData = response.data
            }).catch(error => {
                this.errorMsg = 'Error retrieving search results.'
                this.tableData = []
            })
        },
        validateSingleSearch() {
            switch (this.ballLabel) {
                case 'Mega Ball':
                    if (this.value1.getDay() === 2 || this.value1.getDay() === 5) {
                        this.alertShow = false
                        this.searchSingle()
                    } else {
                        this.alertShow = true
                        this.tableData = []
                    }
                    break
                case 'Powerball':
                    if (this.value1.getDay() === 3 || this.value1.getDay() === 6) {
                        this.alertShow = false
                        this.searchSingle()
                    } else {
                        this.alertShow = true
                        this.tableData = []
                    }
                    break
            }
        }
    }
}
</script>

<style>
    .el-select, .el-date-picker, .el-button {
        margin: 3px 0px 3px 0px;
    }
    .el-alert {
        margin: 3px;
        padding: 3px;
    }
</style>