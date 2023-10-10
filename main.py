import pandas as pd
from sklearn.tree import DecisionTreeClassifier  # Import Decision Tree Classifier
# Import train_test_split function
from sklearn.model_selection import train_test_split
# Import scikit-learn metrics module for accuracy calculation
from sklearn import metrics


df = pd.read_csv("assets/surv.csv")

print(df.dropna(inplace=True))

df["Sex"].replace("male", 1, inplace=True)
df["Sex"].replace("female", 2, inplace=True)

df["Embarked"].replace("S", 1, inplace=True)
df["Embarked"].replace("C", 2, inplace=True)
df["Embarked"].replace("Q", 3, inplace=True)


print(df.isna().sum())


X = df[["Pclass", "Sex", "Age", "SibSp", "Parch", "Fare", "Embarked"]]

y = df.Survived

# print(X)

X_train, X_test, y_train, y_test = train_test_split(
    X, y, test_size=0.5, random_state=1)

clf = DecisionTreeClassifier(
    max_depth=5, criterion="entropy", min_samples_split=50)

# Train Decision Tree Classifer
clf = clf.fit(X_train, y_train)

# Predict the response for test dataset
y_pred = clf.predict(X_test)

print("Accuracy:", metrics.accuracy_score(y_test, y_pred))
