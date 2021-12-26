<template>
  <div>
    <div id="above">
      <table width="90%" class="table">
        <caption>
          <h2>共有 {{counter}} 台主机</h2>
        </caption>
        <thead>
          <tr>
            <th>设备识别ID</th>
            <th>设备名称</th>
            <th>最后更新时间</th>
            <th>初次上线时间</th>
      <!--      <th>详细记录</th>-->
            <th>详细记录</th>
          </tr>
        </thead>
        <tr v-for="beacon in beacons" :key="beacon.DeviceID">
          <td>{{beacon.DeviceID}}</td>
          <td>{{beacon.DeviceName}}</td>
          <td>{{getDate(beacon.LastUpdateTime)}}</td>
          <td>{{getDate(beacon.RegisterTime)}}</td>
    <!--      <td><button @click="gotoDetails(beacon.DeviceID)">详细</button></td>-->
          <td><router-link v-bind:to="'/'+beacon.DeviceID" tag="button">详细</router-link></td>
        </tr>
      </table>
    </div>
    <br>
    <hr>
    <div id="below">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "show-beacons",
  data(){
    return{
      beacons: null,
      counter: 0,
      deviceID: null
    }
  },
  created() {
    axios.get("http://42.193.116.23:4435/v1/beacons")
    // axios.get("http://localhost:4435/v1/beacons")
        .then(response => {
          // console.log(response.data)
          this.beacons = response.data;
          this.counter = response.data.length;
        });
  },
  methods:{
    getDate: function(timestamp) {
      var realTimestamp = timestamp*1000
      var date = new Date(realTimestamp);
      var Y = date.getFullYear() + '-';
      var M = (date.getMonth()+1 < 10 ? '0'+(date.getMonth()+1) : date.getMonth()+1) + '-';
      var D = date.getDate() + ' ';
      var h = date.getHours() + ':';
      var m = date.getMinutes() + ':';
      var s = date.getSeconds();
      return Y+M+D+h+m+s
    },
    gotoDetails: function (deviceID) {
      // 转向详细界面
      console.log(deviceID)
      history.pushState({},'',deviceID)//前端路由
      // history.back() // 返回的时候用
    }
  }
}
</script>

<style scoped>
hr
{
  border: 3px solid #CCE8EB;
  color: #666;
  border-radius: 5px;
}
table
{
  border-collapse: collapse;
  margin: 0 auto;
  text-align: center;
  width: 80%;
}
table td, table th
{
  border: 1px solid #cad9ea;
  color: #666;
  height: 30px;
}
table thead th
{
  background-color: #CCE8EB;
}
table tr:nth-child(odd)
{
  background: #fff;
}
table tr:nth-child(even)
{
  background: #F5FAFA;
}
</style>