#!/bin/bash
exclude='github'
files=`ls | grep -v $exclude`
for each_file in $files
do
cp -r $each_file github.com/dszhengyu/launch/kid/
done
git config --global user.email "dszhengyu@qq.com"
git config --global user.name "appendix"
