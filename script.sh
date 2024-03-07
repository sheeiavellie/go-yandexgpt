git log --all --oneline -- storyteller-api/pkg | while read line;
do     
    hash=$(echo $line | cut -d ' ' -f 1);     
    git cherry-pick $hash
    git cherry-pick --continue
done
