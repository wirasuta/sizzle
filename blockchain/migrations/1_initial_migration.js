const Sizzle = artifacts.require("Sizzle");
const EnumerableSet = artifacts.require("EnumerableSet");

module.exports = function (deployer) {
  deployer.deploy(EnumerableSet);
  deployer.link(EnumerableSet, Sizzle);
  deployer.deploy(Sizzle);
};
