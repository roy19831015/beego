echo "git commit -a -m '第v2.0.$1版本'"
git commit -a -m "第v2.0.$1版本"
echo "git checkout -b release/v2.0.$1"
git checkout -b release/v2.0.$1
echo "git push -u origin release/v2.0.$1"
git push -u origin release/v2.0.$1
echo "git tag v2.0.$1"
git tag v2.0.$1
echo "git push --tags"
git push --tags