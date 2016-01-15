#! /bin/bash
case "$1" in
  run) shift; lein fregec :run Diamond ;;
  test) shift; lein fregec :test DiamondTest ;;
  *) lein fregec :run Diamond
esac
