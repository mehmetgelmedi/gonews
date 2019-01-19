<template>
<div>
    <div class="bg"></div>
    <div class="wrap">
    <div class="square">
        <h1><a href="https://newsapi.org/">Thanks</a><br/>
        <div class="underline">z</div>
        </h1>
    </div>
    </div>
    </div>
</template>

<script>
export default {
    name:"NewsAPI"
}
</script>

<style lang="sass" scoped>
////thx to https://codepen.io/kristarling/pen/eZQZrL
@font-face
  font-family: thunder
  src: url(https://s3-us-west-2.amazonaws.com/s.cdpn.io/507450/thunder-z.ttf)

@font-face
  font-family: againts
  src: url(https://s3-us-west-2.amazonaws.com/s.cdpn.io/507450/againts.otf)
  
$bg: #FF4173
$bg1: #FFCE49
$white: #FFF

body
  background: $bg
  color: $white
  text-align: center
  overflow: hidden
  
.bg
  position: absolute
  top: 50vh
  left: 40vw
  margin: -400px 0 0 0
  z-index: -1
  height: 200vh
  width: 200vw
  background-size: 600px
  opacity: 0.1
  
.wrap
  position: relative
  top: 50vh
  margin-top: -280px
  
.square
  background: $bg1
  width: 258px
  height: 280px
  border-radius: 12px
  overflow: visible
  position: relative
  margin: -140px auto 0
  box-shadow: 0 4px 25px rgba(0,0,0,0.28)
  z-index: 10

h1
  position: relative
  left: -54px
  top: 20px
  font-family: againts, cursive
  font-weight: normal
  font-size: 11em
  transform: rotate(-10deg)
  text-shadow: 0 3px 15px rgba(0,0,0,0.2)
  letter-spacing: 3px
  
  .underline
    right: -112px
    position: relative
    bottom: 130px
    font-size: 1.35em
    font-family: thunder
    transform: rotate(180deg)
    transform: scale(-1, 1)
    
.line
  margin: 30px auto 6px
  background: $white
  width: 50px
  height: 2px
  
.subtext
  font-size: 1.2em
  color: rgba($white, 0.35)
  font-weight: bold
</style>

