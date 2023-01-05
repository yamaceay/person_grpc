for FILE in deploy/*; do
  kubectl apply -f $FILE
done