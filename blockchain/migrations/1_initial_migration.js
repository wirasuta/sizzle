const Sizzle = artifacts.require("Sizzle");
const EnumerableSet = artifacts.require("EnumerableSet");
const SignedMath = artifacts.require("SignedMath");

module.exports = function (deployer) {
  deployer.deploy(EnumerableSet);
  deployer.deploy(SignedMath);
  deployer.link(EnumerableSet, Sizzle);
  deployer.link(SignedMath, Sizzle);
  deployer.deploy(Sizzle);
};
