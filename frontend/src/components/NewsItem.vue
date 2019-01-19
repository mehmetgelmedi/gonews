<template>
    <div class="NewsItem card">
        <div class="wrapper" :style="{'background': 'url('+n.urlToImage+') center / cover no-repeat'}">
            <div class="header">
                <div class="date">
                    <span>{{n.publishedAt}}</span>
                </div>
            </div>
            <div class="data">
                <div class="content">
                    <span class="author">{{n.author}}</span>
                    <h1 class="title">
                        <a target="_blank" v-bind:href="n.url">{{n.title}}</a>
                    </h1>
                    <p class="text">{{n.content}}</p>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: "NewsItem",
        props:["n"]
    };
</script>

<style lang="scss" scoped>
//thx to https://codepen.io/choogoor/pen/YWBxAg
$bermuda: #77d7b9;
$white: #fff;
$clay:#813F0B;
$orange: #FD6A02;
$melon: #F79862;
$gold:#F9A602;

$black: #000;
// clear-fix mixin
@mixin cf {
  &::before,
  &::after {
    content: '';
    display: table;
  }
  &::after {
    clear: both;
  }
}
// Base styles
.card {
  float: left;
  padding: 5px;
  width: 50%;
  .menu-content {
    @include cf;
    margin: 0;
    padding: 0;
    list-style-type: none;
    li {
      display: inline-block;
    }
    a {
      color: $white;
    }
    span {
      position: absolute;
      left: 50%;
      top: 0;
      font-size: 10px;
      font-weight: 700;
      font-family: 'Open Sans';
      transform: translate(-50%, 0);
    }
  }
  .wrapper {
    background-color: $white;
    min-height: 540px;
    position: relative;
    overflow: hidden;
    box-shadow: 0 19px 38px rgba($black, 0.3), 0 15px 12px rgba($black, 0.2);
    &:hover {
      .data {
        transform: translateY(0);
      }
    }
  }
  .data {
    position: absolute;
    bottom: 0;
    width: 100%;
    transform: translateY(calc(70px + 1em));
    transition: transform 0.3s;
    .content {
      padding: 1em;
      position: relative;
      z-index: 1;
    }
  }
  .author {
    text-shadow: 2px 2px 5px $gold;
    font-size: 12px;
  }
  .title {
    margin-top: 10px;
    text-shadow: 2px 2px 5px $black;
    font-weight: bold;
  }
  .text {
      color:$black;
      font-weight: bold;
      text-shadow: 2px 2px 5px $white;
    height: 70px;
    margin: 0;
  }
  input[type='checkbox'] {
    display: none;
  }
  input[type='checkbox']:checked + .menu-content {
    transform: translateY(-60px);
  }
}
    .NewsItem {
        .wrapper {
            &:hover {
                .menu-content {
                    span {
                        transform: translate(-50%, -10px);
                        opacity: 1;
                    }
                }
            }
        }

        .header {
            @include cf;
            color: $black;
            padding: 1em;
            text-shadow: 2px 2px 5px $gold;

            .date {
                float: left;
                font-size: 12px;
            }
        }

        .menu-content {
            float: right;

            li {
                margin: 0 5px;
                position: relative;
            }

            span {
                transition: all 0.3s;
                opacity: 0;
            }
        }

        .data {
            color: $white;
            transform: translateY(calc(70px + 4em));
        }

        .title {
            a {
                color: $white;
            }
        }

        .button {
            display: block;
            width: 100px;
            margin: 2em auto 1em;
            text-align: center;
            font-size: 12px;
            color: $white;
            line-height: 1;
            position: relative;
            font-weight: 700;

            &::after {
                content: "\2192";
                opacity: 0;
                position: absolute;
                right: 0;
                top: 50%;
                transform: translate(0, -50%);
                transition: all 0.3s;
            }

            &:hover {
                &::after {
                    transform: translate(5px, -50%);
                    opacity: 1;
                }
            }
        }
    }
</style>