<template>
  <el-row>
    <el-collapse accordion>
      <el-collapse-item title="Search Winners" name="1">
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
            <el-button type="primary" icon="search" @click="searchSingle" v-show="showSingleCalendar">Search</el-button>
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
                    label="Date">
                </el-table-column>
                <el-table-column
                    prop="winning_numbers"
                    label="Numbers">
                </el-table-column>
                <el-table-column
                    :prop="ballProp"
                    :label="ballLabel">
                </el-table-column>
                <el-table-column
                    prop="multiplier"
                    :label="multiplier">
                </el-table-column>
            </el-table>
          </div>
        </el-col>
      </el-collapse-item>
    </el-collapse>
  </el-row>

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
            errorMsg: ''
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
        }
    }
}
</script>

<style>
    .el-select, .el-date-picker, .el-button {
        margin: 3px 0px 3px 0px;
    }
</style>